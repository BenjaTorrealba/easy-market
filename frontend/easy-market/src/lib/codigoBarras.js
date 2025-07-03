import Quagga from 'quagga';

let detectedCallback = null;
let quaggaIniciado = false; // <--- NUEVO
export function iniciarScanner(onDetected) {
    // Elimina el listener anterior si existe
    if (detectedCallback) {
        Quagga.offDetected(detectedCallback);
    }

    detectedCallback = result => {
        if (onDetected) {
            onDetected(result.codeResult.code);
        }
        Quagga.stop();
        Quagga.offDetected(detectedCallback); // Limpia el listener después de usarlo
        detectedCallback = null;
    };

    Quagga.init({
        inputStream: {
            type: "LiveStream",
            target: document.querySelector('#scanner-container'),
            constraints: {
                facingMode: "environment"
            }
        },
        decoder: {
            readers: [
                "ean_reader",
            ] 
        }
    }, function(err) {
        if (err) {
            console.error(err);
            return;
        }
        Quagga.start();
        quaggaIniciado = true; // <--- SOLO si inicia bien
        Quagga.onDetected(detectedCallback);
    });
}


export function detenerScanner() {
    if (quaggaIniciado) { // <--- SOLO si está iniciado
        Quagga.stop();
        quaggaIniciado = false;
    }
    if (detectedCallback) {
        Quagga.offDetected(detectedCallback);
        detectedCallback = null;
    }
}