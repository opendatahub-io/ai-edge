# Tensorflow Face Detection
## Description
Face detector based on SqueezeNet light (half-channels) as a backbone with a single SSD for indoor/outdoor scenes shot by a front-facing camera. The backbone consists of fire modules to reduce the number of computations. The single SSD head from 1/16 scale feature map has nine clustered prior boxes.

## Test
### Run OVMS with the model inside
Execute the following command from the [pipelines](../../) folder:
```
docker run -d -u $(id -u):$(id -g) --rm -v ${PWD}/models:/model:Z -p 9000:9000 openvino/model_server:latest --model_name face-detection --model_path /model/tensorflow-facedetection --port 9000 --shape auto
```

### Test call

```
git clone https://github.com/openvinotoolkit/model_server.git
cd model_server/demos/face_detection/python
pip install -r ../../common/python/requirements.txt
# In case of errors remove the tensorflow-serving-api version from the ../../common/python/requirements.txt
mkdir results
python face_detection.py --batch_size 1 --width 300 --height 300 --grpc_port 9000
# Open the results folder
```
## Credits
- https://docs.openvino.ai/archive/2022.1/omz_models_model_face_detection_retail_0004.html
- https://docs.openvino.ai/2023.2/ovms_demo_face_detection.html
