from flask import Flask

app = Flask(__name__)

# This line adds the hasura example routes form the hasura.py file.
# Delete these two lines, and delete the file to remove them from your project
from .hasura import hasura_examples
app.register_blueprint(hasura_examples)

from .server import *
