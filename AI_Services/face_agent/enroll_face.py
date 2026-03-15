import psycopg2
from pgvector.psycopg2 import register_vector
from deepface import DeepFace

conn = psycopg2.connect("host=localhost dbname=sentinel user=postgres password=postgres port=5432")
register_vector(conn)
cur=conn.cursor()

def enroll(name,img_path):
    print(f"Extracting embedding for {name}")
    results=DeepFace.represent(img_path=img_path,model_name="Facenet512")
    embedding = results[0]["embedding"]

    cur.execute("INSERT INTO faces (names,embedding) VALUES (%s,%s)",(name,embedding))
    conn.commit()

    print("successfully enrolled {name}")

if __name__=="__main__":
    enroll("Phani","myface_jpg")