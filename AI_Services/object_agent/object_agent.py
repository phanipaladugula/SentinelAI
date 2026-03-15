import os
import time
from ultralytics import YOLO

model = YOLO('yolov8n.pt')

FRAME_PATH = r"C:\Users\phani\VScode\Projects\SentinelAI\Backend"

last_alert_time={}
ALERT_COOLDOWN = 30 

def run_agent():
    print("ai object agent active. monitoring frames")

    while True:
        frames=[f for f in os.listdir(FRAME_PATH) if f.startswith("camera_") and f.endswith(".jpg") ]

        for frame in frames:
            camera_id = frame.split('_')[1].split('.'[0])
            full_path=os.path.join(FRAME_PATH,frame)

            results=model(full_path,conf=0.5,verbose=False)

            person_detected=False
            for r in results:
                for box in r.boxes:
                    label=model.names[int(box.cls[0])]
                    if label=='person':
                        person_detected=True
                        break
            if person_detected:
                current_time=time.time()

                if camera_id not in last_alert_time or (current_time - last_alert_time[camera_id])>ALERT_COOLDOWN:
                    print(f"[ALERT] {time.strftime('%H:%M:%S')} - Intrusion detected on Camera: "{camera_id})
                    last_alert_time[camera_id]=current_time
        time.sleep(1)
        time.sleep(1)
if __name__=="__main__":
    run_agent()