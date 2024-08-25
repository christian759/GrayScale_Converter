from PIL import Image
import os

folder_path: str = "picture_folder"
images: list = []
imagePaths: list[str] = []

for img in os.listdir(folder_path):
    imagepath = os.path.join(folder_path, img)
    image = Image.open(imagepath)
    images.append(image)
    imagePaths.append(imagepath)

j: int = 0
print(len(imagePaths))
for i in range(0, len(imagePaths)):
    image = Image.open(imagePaths[i])
    image = image.convert("L")
    image.save("grayscale_image"+str(i)+".jpg")
