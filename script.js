const go = new Go()

WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(result => {
    go.run(result.instance)

    document.getElementById("explain-input").addEventListener("keydown", (ev) => {
        if (ev.key === 'Enter' || ev.key === 'Return') {
            document.getElementById("explain-button").click();
        }
    })

    document.getElementById("evaluate-button").addEventListener("click", () => {
        let input = document.getElementById("explain-input").value
        let resstr = evaluate(input)
        let res = JSON.parse(resstr)

        if (res.err){
            document.getElementById("result").innerHTML = res.err
            return
        }

        document.getElementById("result").innerHTML = "<div>" + res.value + "</div>" +  "<hr/>" + "<br/>"
    });

    document.getElementById("explain-button").addEventListener("click", () => {
        let input = document.getElementById("explain-input").value
        let resstr = explain(input)
        let res = JSON.parse(resstr)
        let result = ""
    
        if (res.err){
            document.getElementById("result").innerHTML = res.err
            return
        }
    
        res.value.forEach(step => {
            result += "<div>"
            step.EquivalentForms.forEach(equivalentForm => {
                result += equivalentForm + "<i>→</i>"
            })
            result += step.Result + "</div>" + "<hr/>" + "<br/>"
        });
    
        document.getElementById("result").innerHTML = result
    });
});