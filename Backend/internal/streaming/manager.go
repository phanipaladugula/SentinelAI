package streaming

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sentinel-ai/internal/camera"
	"sentinel-ai/internal/logger"
	"sync"
	"time"

	"go.uber.org/zap"
)

type StreamManager struct{
	repo *camera.Repository
	processes map[string]*exec.Cmd
	mu sync.Mutex
}

func NewStreamManager(repo *camera.Repository)*StreamManager{
	return &StreamManager{
		repo:repo,
		processes: make(map[string]*exec.Cmd),
	}
}

func(m *StreamManager)StartIngestion(ctx context.Context){
	ticker:=time.NewTicker(10*time.Second)
	defer ticker.Stop()
	
	logger.Log.Info("Stream Manager started")

	for{
		select{
		case <-ctx.Done():
			logger.Log.Info("Stream Manager stopping...")
			return
		case <-ticker.C:
			cameras,err:=m.repo.GetAll(ctx)
			if err!=nil{
				logger.Log.Error("Failed to fetch cameras for ingestion",zap.Error(err))
				continue
			}

			for _,cam:=range cameras{
				if cam.Status=="online"{
					go m.startFFmpeg(ctx,cam)
				}
			}
		}
	}
}

func(m *StreamManager)startFFmpeg(ctx context.Context,cam camera.Camera){
	m.mu.Lock()
	if _,exists:=m.processes[cam.ID];exists{
		m.mu.Unlock()
		return
	}
	m.mu.Unlock()

	wd,_:=os.Getwd()
	outputPath := filepath.Join(wd,fmt.Sprintf("camera_%s.jpg", cam.ID))

	cmd:=exec.CommandContext(ctx,"ffmpeg",
		"-re",
		"-stream_loop","-1",
		"-i",cam.RTSPUrl,
		"-vf","fps=1",
		"-update","1",
		"-y",
		"-an",
		outputPath,
	)
	cmd.Stdout=os.Stdout
	cmd.Stderr=os.Stderr

	m.mu.Lock()
	m.processes[cam.ID]=cmd
	m.mu.Unlock()

	logger.Log.Info("Starting stream ingestion", zap.String("camera",cam.Name))

	if err:=cmd.Run();err!=nil{
		if ctx.Err()==nil{
		 logger.Log.Error("FFmpeg process exited", zap.String("camera",cam.Name),zap.Error(err))
		}
		m.mu.Lock()
		delete(m.processes,cam.ID)
		m.mu.Unlock()
	}
}