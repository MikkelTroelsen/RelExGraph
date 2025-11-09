from relik import Relik
from relik.inference.data.objects import RelikOutput, Triplets
from flask import Flask, request, jsonify

app = Flask(__name__)

relik = Relik.from_pretrained(
    "relik-ie/relik-relation-extraction-small",
)

@app.route('/get-relations', methods=['POST'])
def get_relations():
    try:
        data = request.get_json()
        print(data)
        text = data.get('text') if data else None
        out: RelikOutput = relik(text)
        triples = extract_triples(out.triplets)
        return jsonify(triples)
    except:
        return jsonify(error="internal error"), 500

def extract_triples(triples: list[Triplets]):
    return [(triple.subject.text, triple.label, triple.object.text) for triple in triples]

if __name__ == "__main__":
    app.run()
