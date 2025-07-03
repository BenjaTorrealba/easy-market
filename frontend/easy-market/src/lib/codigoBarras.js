import Quagga from 'quagga';

export function iniciarScanner(onDetected) {
    Quagga.init({
        inputStream: {
            type: "LiveStream",
            target: document.querySelector('#scanner-container'), // Asegúrate de tener este div en tu HTML
            constraints: {
                facingMode: "environment" // Usa la cámara trasera si está disponible
            }
        },
        decoder: {
            readers: [
                "ean_reader",
                "ean_8_reader",
                "code_128_reader",
                "code_39_reader",
                "code_39_vin_reader",
                "upc_reader",
                "upc_e_reader",
                "codabar_reader",
                "i2of5_reader",
                "2of5_reader",
                "code_93_reader"
            ] // Puedes agregar más tipos si lo necesitas
        }
    }, function(err) {
        if (err) {
            console.error(err);
            return;
        }
        Quagga.start();
    });

    Quagga.onDetected(result => {
        if (onDetected) {
            onDetected(result.codeResult.code);
        }
        // Si quieres detener el escaneo después de detectar uno, descomenta la siguiente línea:
        // Quagga.stop();
    });
}

export function detenerScanner() {
    Quagga.stop();
}