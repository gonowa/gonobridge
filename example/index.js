const gonobridge = require("gonobridge");

let goEmitter = gonobridge.Init("go.wasm", ["test", "sdsd"]);

goEmitter.on("hello", function (v) {
    console.log("recieved", v)
});

goEmitter.on("callback", function (log) {
    console.log(log);
    goEmitter.emit("test", 33);
    setTimeout(function () {
        console.log("quitting");
        goEmitter.quit()
    }, 3000)
});


