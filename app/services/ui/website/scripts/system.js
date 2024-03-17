const sdk = new MetaMaskSDK.MetaMaskSDK({
    dappMetadata: {
      name: "Pure JS example",
      url: window.location.host,
    },
    logging: {
      sdk: false,
    }
});

$.ajaxSetup({
    contentType: "application/json; charset=utf-8",
});

function hexer(input) {
    const utf8 = toUTF8Array(input);
    const hex = utf8.map(n => n.toString(16));
    return '0x' + hex.join('');
}

function toUTF8Array(str) {
    var utf8 = [];

    for (var i=0; i < str.length; i++) {
        var charcode = str.charCodeAt(i);

        if (charcode < 0x80) {
            utf8.push(charcode);
        }
        else if (charcode < 0x800) {
            utf8.push(0xc0 | (charcode >> 6),
                    0x80 | (charcode & 0x3f));
        }
        else if (charcode < 0xd800 || charcode >= 0xe000) {
            utf8.push(0xe0 | (charcode >> 12),
                    0x80 | ((charcode>>6) & 0x3f),
                    0x80 | (charcode & 0x3f));
        }
        else {
            // Surrogate pair.
            i++;
            
            // UTF-16 encodes 0x10000-0x10FFFF by
            // subtracting 0x10000 and splitting the
            // 20 bits of 0x0-0xFFFFF into two halves
            charcode = 0x10000 + (((charcode & 0x3ff)<<10)
                    | (str.charCodeAt(i) & 0x3ff));
            utf8.push(0xf0 | (charcode >>18),
                    0x80 | ((charcode>>12) & 0x3f),
                    0x80 | ((charcode>>6) & 0x3f),
                    0x80 | (charcode & 0x3f));
        }
    }

    return utf8;
}

function currentDateTime() {
    const dt = new Date();
    
    const year    = dt.getUTCFullYear();
    const month   = String(dt.getUTCMonth() + 1).padStart(2, '0'); // Month (0-indexed)
    const day     = String(dt.getUTCDate()).padStart(2, '0');
    const hours   = String(dt.getUTCHours()).padStart(2, '0');
    const minutes = String(dt.getUTCMinutes()).padStart(2, '0');
    const seconds = String(dt.getUTCSeconds()).padStart(2, '0');

    return `${year}${month}${day}${hours}${minutes}${seconds}`;
}