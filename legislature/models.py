from flask.ext.sqlalchemy import SQLAlchemy

from legislature import app

db = SQLAlchemy(app)
	

class ModelUtils:
	""" Provide utility functions for saving and deleting individual entities
	stored in the SQLAlchemy database.
	"""
	def put(self):
		db.session.add(self)
		db.session.commit()

	def delete(self):
		db.session.delete(self)
		db.session.commit()


class Tag(db.Model, ModelUtils):
	""" A tag or category for a bill.
	"""
	id = db.Column(db.Integer, primary_key=True)
	name = db.Column(db.String(255))

	def __repr__(self):
		return '<Tag %r>' % self.name

	def __str__(self):
		return self.name


class Bill(db.Model, ModelUtils):
	""" Store data about a bill.
	"""
	id = db.Column(db.Integer, primary_key=True)
	title = db.Column(db.Text)
	description = db.Column(db.Text)
	db.relationship('Tag', backref='bill', lazy='dynamic')

	timestamp = db.Column(db.DateTime)

	def __repr__(self):
		return '<Bill %r>' % self.title

	def __str__(self):
		return self.title
