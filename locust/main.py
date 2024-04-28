from locust import HttpUser,task,between
import json
class Lectura():
    def __init__(self) -> None:
        self.array=[]
    
    def elegirRandom(self):
        if len(self.array==0):
    
            pass
    
    def cargar(self):
        try:
            with open("data.json","r") as f:
                self.array=json.loads(f.read())
        except:
            print("No se pudo cargar")
    
    def sacar(self):
        if len(self.array)>0:
            return self.array.pop()
        return None

class  QuickStart(HttpUser):
    wait_time= between(0.5,1)
    lect = Lectura()
    lect.cargar()

    @task
    def send_data(self):
        dato=self.lect.sacar()
        if dato!=None:
            self.client.post("/",json=dato)
        else:
            print("Envio finalizado") 
            self.stop(True)

        ##print("Ya no hay datos")
        
        
    def on_start(self) -> None:
        print("Iniciando carga")