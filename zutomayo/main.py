from flask import Flask, render_template

app = Flask(__name__)

@app.route("/")
def hello_world():
    return "<p>Hello, World!</p>"

@app.get("/halo")
def hello():
    return render_template("hello.html",person="name")

app.run("localhost",4000)
