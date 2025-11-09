from relik import Relik
from relik.inference.data.objects import RelikOutput, Triplets
from flask import Flask, request, jsonify

app = Flask(__name__)

relik = Relik.from_pretrained(
        "relik-ie/relik-relation-extraction-small",
    )

@app.route('get-relations', metholds=['GET'])
def get_relations():
    try:
        extraction_text = request.args.get('text', type=str)
        out: RelikOutput = relik(extraction_text)
        triples = extract_triples(out.triplets)
        return jsonify(triples)
    except:
        return jsonify(error="internal error"), 500

def extract_triples(triples: list[Triplets]):
    return [(triple.subject.text, triple.label, triple.object.text) for triple in triples]

if __name__ == "__main__":
    app.run()
