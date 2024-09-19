import datetime

from fastapi import FastAPI, Depends
from fastapi import Request, Form
from fastapi.responses import HTMLResponse
from fastapi.templating import Jinja2Templates
from sqlalchemy import create_engine
from sqlalchemy.orm import Session
from sqlalchemy.orm import sessionmaker
from fastapi.staticfiles import StaticFiles

from models.models import Recipe
from models.models import Base





# строка подключения
SQLALCHEMY_DATABASE_URL = "sqlite:///./sql_app.db"

# создаем движок SqlAlchemy
engine = create_engine(
    SQLALCHEMY_DATABASE_URL, connect_args={"check_same_thread": False}
)

# создаем таблицы
Base.metadata.create_all(bind=engine)

# создаем сессию подключения к бд
SessionLocal = sessionmaker(autoflush=False, bind=engine)

app = FastAPI(
    title='SalaryApp'
)







def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()




templates = Jinja2Templates(directory="templates")



@app.get("/", response_class=HTMLResponse)
async def get_specific_operations(request: Request, db: Session = Depends(get_db)):
    return templates.TemplateResponse("index.html", {"request": request, 'kek': db.query(Recipe).all()})






@app.post("/postdata")
async def postdata(db: Session = Depends(get_db)):
    pass
    


