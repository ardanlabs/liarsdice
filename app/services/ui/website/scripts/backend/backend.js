import App from './app.js';

const app = new App('http://0.0.0.0:3000');

const sdk = new MetaMaskSDK.MetaMaskSDK({
    dappMetadata: {
        name: 'Pure JS example',
        url: window.location.host,
    },
    logging: {
        sdk: false,
    },
});

$.ajaxSetup({
    contentType: 'application/json; charset=utf-8',
});

window.onload = function () {
    app.Init();
};
