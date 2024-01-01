const go = new Go()

WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(wasmModule => {
    go.run(wasmModule.instance)

    
    const textEnc = new TextEncoder()
    const textDec = new TextDecoder()
    const linierMemory = new Uint8Array(wasmModule.instance.exports.memory.buffer)
    
    function malloc(value) {
        const input = textEnc.encode(value)
        const ptr = wasmModule.instance.exports.lmalloc(input.length)
        for (let i = 0; i < input.length; i++) {
            linierMemory[ptr + i] = input[i]
        }
    }

    function parseResult(ptr) {
        const size = wasmModule.instance.exports.size()
        const str = textDec.decode(linierMemory.slice(ptr, ptr + size))
        return JSON.parse(str)
    }

    document.getElementById("explain-input").addEventListener("keydown", (ev) => {
        if (ev.key === 'Enter' || ev.key === 'Return') {
            document.getElementById("explain-button").click();
        }
    })

    document.getElementById("evaluate-button").addEventListener("click", () => {
        let value = document.getElementById("explain-input").value
        malloc(value)

        const res = parseResult(wasmModule.instance.exports.evaluate())
        if (res.err){
            document.getElementById("result").innerHTML = res.err
            return
        }

        document.getElementById("result").innerHTML = "<div>" + res.value + "</div>" +  "<hr/>" + "<br/>"
    });
    
    document.getElementById("explain-button").addEventListener("click", () => {
        let value = document.getElementById("explain-input").value
        malloc(value)

        const res = parseResult(wasmModule.instance.exports.explain())
        if (res.err){
            document.getElementById("result").innerHTML = res.err
            return
        }
        
        let result = ""
        res.value.forEach(step => {
            result += "<div>"
            step.EquivalentForms?.forEach((equivalentForm, i) => {
                result += equivalentForm
                result += "<i>→</i>"
            })
            result += "<strong>" + step.Result + "</strong>"
            result += "</div>"

            if(step.Explaination) {
                result += "<div class=\"explanation\">"+
                               "<p>"+
                                    "<small style=\"font-weight:500\">Bitwise Operation Explanation:\n</small>"+
                                    step.Explaination+
                                    "&nbsp;&nbsp;≈&nbsp;&nbsp;"+
                                    "<strong>"+step.Result+"</strong>"+
                                "</p>"+
                           "</div>"
            }

            result += "<hr/>"
        });
    
        document.getElementById("result").innerHTML = result
    });
});