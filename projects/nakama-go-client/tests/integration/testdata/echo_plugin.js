function rpcEcho(ctx, logger, nk, payload) {
    return payload;
}

function rpcUppercase(ctx, logger, nk, payload) {
    var data = JSON.parse(payload);
    return JSON.stringify({ result: String(data.text).toUpperCase() });
}

function rpcFail(ctx, logger, nk, payload) {
    throw Error("intentional failure");
}

function InitModule(ctx, logger, nk, initializer) {
    initializer.registerRpc("echo", rpcEcho);
    initializer.registerRpc("uppercase", rpcUppercase);
    initializer.registerRpc("fail", rpcFail);
}
