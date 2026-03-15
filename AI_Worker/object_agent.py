import os
import time
from ultralyics import YOLO

model = YOLO('yolov8n.pt')

FRAME_PATH = r"C:\Users\phani\VScode\Projects\SentinelAI\Backend"

def run_agent():
    print("ai object agent active. monitoring frames")

    while True:
        frames=[f for f in os.listdir(FRAME_PATH) if f.startswith("camera_") and f.endswith(".jpg") ]

        for frame in frames:
            full_path=os.path.join(FRAME_PATH,frame)

            results=model(full_path,conf=0.5,verbose=False)

            for r in results:
                for box in r.boxes:
                    label=model.names[int(box.cls[0])]
                    print(f"[{time.strftime('%H:%M:%S')}] {frame} -> Detected: {label}")

                    if label=='person':
                        print("Intrusion alert: person detected!")
        time.sleep(1)
if __name__=="__main__":
    run_agent()