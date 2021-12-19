//MJJ MUSIC
'use strict';

function requestGC(onSuccess, onFailed) {
    $.ajax({
        url: '/manipulate/gc',
        type: 'POST',
        data: {
            token: tokenProvider.get()
        },
        success: function (data) {
            if (isStatusOK(data.status)) {
                onSuccess();
            } else {
                tokenProvider.clear();
                onFailed(data.status);
            }
        }
    });
}

function requestMemoryInfo(onSuccess, onFailed) {
    $.ajax({
        url: '/manipulate/mem',
        type: 'POST',
        data: {
            token: tokenProvider.get()
        },
        success: function (data) {
            if (isStatusOK(data.status)) {
                onSuccess(data.info);
            } else {
                tokenProvider.clear();
                onFailed(data.status);
            }
        }
    });
}