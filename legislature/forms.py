from flask_wtf import Form
from wtforms import TextField
from wtforms.validators import DataRequired, TextAreaField, Regexp

ValidList = Regexp('(?:\w+)(?:,\w+)*', message="Not a comma-separated list of tags.")

class BillForm(Form):
	name = TextField('name', validators=[DataRequired])
	tags = TextAreaField('tags', validators=[DataRequired, ValidList])