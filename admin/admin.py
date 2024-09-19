from sqladmin import Admin, ModelView
from main import app, engine
from models.models import Recipe

admin = Admin(app, engine)


class UserAdmin(ModelView, model=Recipe):
    column_list = [Recipe.id, Recipe.title]

admin.add_view(UserAdmin)