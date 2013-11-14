from flask import Flask

app = Flask(__name__)

import legislature.models
import legislature.views
