import Quagga from "quagga";

let detectedCallback = null;
let quaggaIniciado = false;
export function iniciarScanner(onDetected) {
  if (detectedCallback) {
    Quagga.offDetected(detectedCallback);
  }

  detectedCallback = (result) => {
    if (onDetected) {
      onDetected(result.codeResult.code);
    }
    Quagga.stop();
    Quagga.offDetected(detectedCallback);
    detectedCallback = null;
  };

  Quagga.init(
    {
      inputStream: {
        type: "LiveStream",
        target: document.querySelector("#scanner-container"),
        constraints: {
          facingMode: "environment",
        },
      },
      decoder: {
        readers: ["ean_reader"],
      },
    },
    function (err) {
      if (err) {
        console.error(err);
        return;
      }
      Quagga.start();
      quaggaIniciado = true;
      Quagga.onDetected(detectedCallback);
    }
  );
}

export function detenerScanner() {
  if (quaggaIniciado) {
    Quagga.stop();
    quaggaIniciado = false;
  }
  if (detectedCallback) {
    Quagga.offDetected(detectedCallback);
    detectedCallback = null;
  }
}
