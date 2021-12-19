//MJJ MUSIC
'use strict';

async function sha256(message) {
    const hashArray = Array.from(
        new Uint8Array(
            await crypto.subtle.digest(
                'SHA-256',
                new TextEncoder().encode(message)
            )
        )
    );
    return hashArray.map(b => b.toString(16).padStart(2, '0')).join('');
}