from flask import Flask, jsonify, render_template, request

app = Flask(__name__)

@app.route('/')
def index():
        return render_template('index.html')

@app.route('/echo/', methods=['GET'])
def echo():
    ret_data = {"value": request.args.get('echoValue')}
    return jsonify(ret_data)

if __name__ == '__main__':
    app.run(port=8080, debug=True)
