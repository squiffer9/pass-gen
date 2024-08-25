from flask import Flask, render_template, request
import subprocess
import os

app = Flask(__name__)

def generate_password(word1, word2):
    # Adjust the path to the executable based on your project structure
    executable_path = os.path.join("cmd", "password-generator", "password-generator")
    result = subprocess.run(
        [executable_path, word1, word2],
        capture_output=True,
        text=True
    )
    return result.stdout.strip()

@app.route('/', methods=['GET', 'POST'])
def index():
    password = None
    if request.method == 'POST':
        word1 = request.form['word1']
        word2 = request.form['word2']
        password = generate_password(word1, word2)
    return render_template('index.html', password=password)

if __name__ == '__main__':
    app.run(debug=True)

