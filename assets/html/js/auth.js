//MJJ MUSIC
'use strict';

const tokenProvider = {
    get: () => $.cookie("token"),
    set: (val) => $.cookie("token", val),
    clear: () => $.cookie("token", null)
}

const isStatusOK = (val) => val === "ok";

function isTokenValid(onSuccess, onFailed) {
    let localToken = tokenProvider.get();
    if (localToken === null || localToken === undefined) {
        return;
    }
    $.ajax({
        url: "/auth/test",
        type: "POST",
        data: {
            token: localToken
        },
        success: function (data) {
            if (isStatusOK(data.status)) {
                onSuccess();
            } else {
                tokenProvider.clear();
                onFailed();
            }
        }
    });
}

async function login(account, password) {
    alert('ddd');
    $.ajax({
        url: "/auth/req",
        type: "POST",
        data: {
            account: await sha256(account),
            password: await sha256(password)
        },
        success: function (data) {
            if (isStatusOK(data.status)) {
                tokenProvider.set(data.token);
                setInterval(function () {
                    redirectParam("/panel", {"token": tokenProvider.get()});
                }, 500);
            }
            mdui.snackbar({
                message: data.message,
                position: "top"
            });
        },
    });
}