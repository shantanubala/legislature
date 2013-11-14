from flask import render_template

import legislature
app = legislature.app

@app.route("/")
def home_page():
	return render_template("home.html")
