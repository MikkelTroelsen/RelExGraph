FROM python:3.11-slim

WORKDIR /app

# Install system dependencies
RUN apt-get update && apt-get install -y \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Copy requirements first for better caching
COPY requirements.txt .

# Install Python dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Copy application code
COPY extraction_service.py .

# Expose port
EXPOSE 5000

# Add health check endpoint support
RUN echo 'from flask import Flask, jsonify\n\
app = Flask(__name__)\n\
@app.route("/health")\n\
def health():\n\
    return jsonify({"status": "healthy"}), 200' > /tmp/health_check.py

# Run the application
CMD ["python", "extraction_service.py"]
