from flask import Flask, render_template, request, jsonify
from flask_cors import CORS  # Add this import
import tensorflow as tf
import numpy as np
from PIL import Image
import io
import base64
import os
import logging

# Set up logging
logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger(__name__)

app = Flask(__name__)
CORS(app, resources={r"/*": {"origins": "*"}})  # Allow all origins for mobile app

# Environment variables for flexibility
MODEL_PATH = os.getenv('MODEL_PATH', 'best_model2.tflite')
PORT = int(os.getenv('PORT', 8080))  # Cloud Run sets PORT environment variable

# Load the TFLite model
try:
    interpreter = tf.lite.Interpreter(model_path=MODEL_PATH)
    interpreter.allocate_tensors()
    input_details = interpreter.get_input_details()
    output_details = interpreter.get_output_details()
    logger.info("Model loaded successfully")
    logger.info(f"Input details: {input_details}")
    logger.info(f"Output details: {output_details}")
except Exception as e:
    logger.error(f"Error loading model: {str(e)}")
    raise

CLASSES = ['Blackheads', 'Cyst', 'Papules', 'Pustules', 'Whiteheads']

RECOMMENDATIONS = {
    'Blackheads': (
        "Cleanse with a salicylic acid cleanser, such as 'CeraVe Renewing SA Cleanser' or 'Acnes Natural Care Oil Control Cleanser', to help clear clogged pores. "
        "Exfoliate 2–3 times weekly with a BHA product like 'Some By Mi AHA BHA PHA 30 Days Miracle Toner' to remove excess sebum and dead skin cells. "
        "Incorporate a retinoid at night, such as 'Avoskin Miraculous Retinol Ampoule', to prevent future blackheads. "
        "Finish with a non-comedogenic moisturizer, such as 'Wardah Acnederm Day Moisturizer', to maintain skin hydration."
    ),
    'Cyst': (
        "Use a gentle hydrating cleanser, like 'Hada Labo Gokujyun Ultimate Moisturizing Face Wash', to prevent irritation. "
        "Spot-treat with benzoyl peroxide, such as 'Benzolac 2.5%', or sulfur-based products like 'JF Sulfur Acne Care' to reduce inflammation and bacteria. "
        "If cystic acne persists, consult a dermatologist for possible oral medications like isotretinoin."
    ),
    'Papules': (
        "Wash your face with a salicylic acid-based foaming cleanser, such as 'The Body Shop Tea Tree Skin Clearing Facial Wash' or 'Acnes Foaming Wash', to reduce inflammation and unclog pores. "
        "Calm redness and strengthen the skin barrier using niacinamide products like 'Somethinc Niacinamide + Moisture Beet Serum'. "
        "Apply a spot treatment with benzoyl peroxide, such as 'Benzolac 2.5%', to target active papules."
    ),
    'Pustules': (
        "Cleanse with a tea tree oil cleanser, like 'The Body Shop Tea Tree Skin Clearing Facial Wash', or a salicylic acid cleanser, such as 'Wardah Acnederm Pure Foaming Cleanser', to reduce bacteria. "
        "Use a benzoyl peroxide product, such as 'Benzolac 2.5%', as a spot treatment to combat pustules. "
        "Avoid squeezing or popping pustules to minimize scarring and further inflammation."
    ),
    'Whiteheads': (
        "Start with a glycolic or salicylic acid cleanser, such as 'Safi White Expert Oil Control & Acne Cleanser', to remove excess oil and debris. "
        "Exfoliate gently with an AHA toner, like 'Avoskin Miraculous Refining Toner', to prevent the buildup of dead skin cells. "
        "Incorporate a retinoid, such as 'Somethinc Level 1% Retinol Serum', to accelerate cell turnover and prevent whiteheads. "
        "Finish with a lightweight moisturizer, such as 'Emina Ms. Pimple Acne Solution Moisturizing Gel', to keep the skin hydrated without clogging pores."
    ),
}

PRODUCT_IMAGES = {
    'Blackheads': {
        'CeraVe Renewing SA Cleanser': 'https://i5.walmartimages.com/asr/4581933e-9a62-436d-bfa2-1aaba1f829fa_1.3ca56085c3b6b266447077f1767357bb.jpeg',
        'Some By Mi AHA BHA PHA 30 Days Miracle Toner': 'https://image.femaledaily.com/dyn/210/images/prod-pics/product_1549596498_Some_By_Mi_800x800.png',
        'Avoskin Miraculous Retinol Ampoule': 'https://images.tokopedia.net/img/cache/200-square/VqbcmM/2024/10/4/576fc209-3a87-43f9-8810-8f3b607a4f13.jpg?ect=4g',
        'Wardah Acnederm Day Moisturizer': 'https://nihonmart.id/pub/media/catalog/product/cache/8a3fa81d90974d9a7299f2eff309979f/a/c/acnederm_day_moisturizer-min.png'
    },
    'Cyst': {
        'Hada Labo Gokujyun Ultimate Moisturizing Face Wash': 'https://images.soco.id/864f16d2-1564-4bb0-8ba5-957c2154b524-image-0-1615971582528',
        'Benzolac 2.5%': 'https://d2qjkwm11akmwu.cloudfront.net/products/1473-1665760701.webp',
        'JF Sulfur Acne Care': 'https://d2qjkwm11akmwu.cloudfront.net/products/247346_2-8-2022_13-40-32-1665856856.webp'
    },
    'Papules': {
        'The Body Shop Tea Tree Skin Clearing Facial Wash': 'https://res.cloudinary.com/dowzkjtns/image/fetch/f_auto,c_limit,w_3840,q_auto/https://assets.thebodyshop.co.id/products/3c0c4ece-0c6f-4768-b6bf-fdbd13c5cc34.jpg',
        'Somethinc Niacinamide + Moisture Beet Serum': 'https://image.femaledaily.com/dyn/210/images/prod-pics/product_1581047610_Niacidinam_800x800.jpg',
        'Benzolac 2.5%': 'https://d2qjkwm11akmwu.cloudfront.net/products/1473-1665760701.webp'
    },
    'Pustules': {
        'The Body Shop Tea Tree Skin Clearing Facial Wash': 'https://res.cloudinary.com/dowzkjtns/image/fetch/f_auto,c_limit,w_3840,q_auto/https://assets.thebodyshop.co.id/products/3c0c4ece-0c6f-4768-b6bf-fdbd13c5cc34.jpg',
        'Wardah Acnederm Pure Foaming Cleanser': 'https://d2jlkn4m127vak.cloudfront.net/medias/products/slides-2-1648519363.webp',
        'Benzolac 2.5%': 'https://d2qjkwm11akmwu.cloudfront.net/products/1473-1665760701.webp'
    },
    'Whiteheads': {
        'Safi White Expert Oil Control & Acne Cleanser': 'https://images.tokopedia.net/img/cache/700/VqbcmM/2021/2/16/8b0fe7c2-60c0-4e14-b889-c594cf19601c.jpg',
        'Avoskin Miraculous Refining Toner': 'https://images.soco.id/3f794f71-4ed0-4262-a13b-cfb82bd280b2-image-0-1627354259406',
        'Somethinc Level 1% Retinol Serum': 'https://images.somethinc.com/uploads/products/thumbs/800x800/SERUM-09.jpg',
        'Emina Ms. Pimple Acne Solution Moisturizing Gel': 'https://s3-ap-southeast-1.amazonaws.com/img-sociolla/img/p/2/2/6/0/2/22602-large_default.jpg'
    }
}


def preprocess_image(image):
    """Preprocess the image to match model input requirements"""
    try:
        if image.mode != 'RGB':
            image = image.convert('RGB')
        image = image.resize((150, 150))
        img_array = tf.keras.preprocessing.image.img_to_array(image)
        img_array = np.expand_dims(img_array, axis=0)
        img_array = tf.keras.applications.efficientnet.preprocess_input(img_array)
        return img_array
    except Exception as e:
        logger.error(f"Error in preprocess_image: {str(e)}")
        raise

def predict_image(image_array):
    """Make prediction using TFLite model"""
    try:
        if image_array.dtype != np.float32:
            image_array = image_array.astype(np.float32)
        interpreter.set_tensor(input_details[0]['index'], image_array)
        interpreter.invoke()
        output_data = interpreter.get_tensor(output_details[0]['index'])
        return output_data[0]
    except Exception as e:
        logger.error(f"Error in predict_image: {str(e)}")
        raise


@app.route('/')
def home():
    return render_template('index.html')

@app.route('/predict', methods=['POST'])
def predict():
    try:
        logger.debug("Received prediction request")

        # Validate request
        if not request.files and not request.json:
            return jsonify({
                'success': False, 
                'error': 'No image provided'
            }), 400

        # Load image from request
        if 'file' in request.files:
            logger.debug("Processing uploaded file")
            file = request.files['file']
            
            # Additional file validation
            if file.filename == '':
                return jsonify({
                    'success': False, 
                    'error': 'No selected file'
                }), 400
            
            image = Image.open(file)
        elif request.json and 'image' in request.json:
            logger.debug("Processing base64 image")
            try:
                # Handle base64 image from mobile app
                image_data = request.json['image']
                # Remove data URL prefix if present
                if image_data.startswith('data:image'):
                    image_data = image_data.split(',')[1]
                
                image = Image.open(io.BytesIO(base64.b64decode(image_data)))
            except Exception as e:
                return jsonify({
                    'success': False, 
                    'error': f'Invalid image data: {str(e)}'
                }), 400
        else:
            return jsonify({
                'success': False, 
                'error': 'No image provided'
            }), 400

        # Rest of your existing prediction logic remains the same...
        processed_image = preprocess_image(image)
        predictions = predict_image(processed_image)

        # Get the top prediction and confidence
        top_prediction_idx = np.argmax(predictions)
        top_prediction = CLASSES[top_prediction_idx]
        confidence = float(predictions[top_prediction_idx])

        # Retrieve the recommendation
        recommendation = RECOMMENDATIONS.get(top_prediction, "No recommendation available")
        
        # Retrieve product images for the top prediction
        product_images = PRODUCT_IMAGES.get(top_prediction, {})

        # Get all predictions sorted by confidence
        all_predictions = [
            {'class': CLASSES[idx], 'confidence': float(pred)}
            for idx, pred in enumerate(predictions)
        ]
        all_predictions = sorted(all_predictions, key=lambda x: x['confidence'], reverse=True)

        # Return the JSON response
        response = {
            'success': True,
            'prediction': top_prediction,
            'confidence': confidence,
            'recommendation': recommendation,
            'product_images': product_images,
            'all_predictions': all_predictions
        }
        return jsonify(response)

    except Exception as e:
        logger.error(f"Error in predict route: {str(e)}", exc_info=True)
        return jsonify({
            'success': False,
            'error': str(e)
        }), 500

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=PORT)
