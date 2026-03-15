import os
import time
import psycopg2
from pgvector.psycopg2 import register_vector
from deepface import DeepFace

FRAME_PATH=r"C:\Users\phani\VScode\Projects\SentinelAI\Backend"
conn = psycopg2.connect("host=localhost dbname=sentinel user=postgres password=postgres port=5432")
register_vector(conn)
cur=conn.cursor()

def run_face_agent():
    print("Face recognition Agent: searching for known faces")

    while True:
        frames=[f for f in os.listdir(FRAME_PATH) if f.startswith("camera") and f.endswith(".jpg")]

        for frame in frames:
            full_path=os.path.join(FRAME_PATH,frame)

            try:
                results=DeepFace.represent(img_path=full_path,model_name='Facenet512',enforce_detection=True)

                for res in results:
                    embedding=res["embedding"]

                    cur.execute("SELECT name,embedding <=> %s as distance FROM faces WHERE (embedding <=> %s) < 0.3 ORDER BY distance LIMIT 1",(embedding,embedding))
                    row=cur.fetchone()

                    if row:
                        print(f"[MATCH] Identified: {row[0]} (Dist: {row[1]:.4f})")
                    else:
                        print(f"[UNKNOWN] Unauthorized person detected on {frame}")
            except Exception:
                pass
        time.sleep(2)
if __name__=="__main__":
    run_face_agent()
