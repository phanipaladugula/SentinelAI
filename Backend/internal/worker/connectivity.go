package worker

import(
	"context"
	"net"
	"net/url"
	"sentinel-ai/internal/camera"
	"time"
	"go.uber.org/zap"
	"sentinel-ai/internal/logger"
)

type ConnectivityWorker struct{
	repo *camera.Repository
}

func NewConnectivityWorker(repo *camera.Repository) *ConnectivityWorker{
	return &ConnectivityWorker{repo:repo}
}

func (w *ConnectivityWorker)Start(ctx context.Context){
	ticker := time.NewTicker(30*time.Second)
	defer ticker.Stop()

	logger.Log.Info("Connectivity worker started")

	for{
		select{
		case <-ctx.Done():
			return
		case <-ticker.C:
			w.checkCameras(ctx)
		}
	}
}

func (w *ConnectivityWorker) checkCameras(ctx context.Context){
	cameras,err:=w.repo.GetAll(ctx)

	if err!=nil{
		logger.Log.Info("Failed to fetch cameras for health check", zap.Error(err))
		return
	}

	for _,cam:=range cameras{
		status:="online"
		if !isReachable(cam.RTSPUrl){
			status="offline"
		}

		_,w.repo.UpdateStatus(ctx,cam.ID,status)
	}
}

func isReachable(rtspURL string)bool{
	u,err:=url.Parse(rtspURL)

	if err!=nil{
		return false
	}

	host:=u.Host
	if !containsPort(host){
		host=host+":554"
	}

	conn,err:=net.DialTimeout("tcp",host,3*time.Second)
	if err!=nil{
		return false
	}
	conn.Close()
	return false
}

func containsPort(host string)bool{
	_,_,err:=net.SplitHostPort(host)
	return err==nil
}