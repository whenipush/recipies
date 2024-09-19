import datetime

from sqlalchemy import Column, DateTime, ForeignKey, String, Integer, Text 
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import relationship


Base = declarative_base()



class Recipe(Base):
    __tablename__ = "recipes"
    
    id = Column(Integer, primary_key=True, index=True)
    title = Column(String, index=True)
    description = Column(Text)
    cooking_time = Column(Integer)  # время в минутах



