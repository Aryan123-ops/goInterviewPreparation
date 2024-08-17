import qrcode

# Define the URL
url = "https://google.com"

# Generate QR code
qr = qrcode.QRCode(
    version=1,
    error_correction=qrcode.constants.ERROR_CORRECT_L,
    box_size=10,
    border=4,
)
qr.add_data(url)
qr.make(fit=True)

# Create an image from the QR code instance
qr_image = qr.make_image(fill='black', back_color='orange')

# Save or display the image
qr_image.save("qr_code.png")  # Save QR code as a PNG file
qr_image.show()  # Display QR code (this might open the image viewer)
