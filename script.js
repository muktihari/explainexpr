const go = new Go()

WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(result => {
    go.run(result.instance)

    document.getElementById("explain-button").addEventListener("click", () => {
        let input = document.getElementById("explain-input").value
        let resstr = explain(input)
        let res = JSON.parse(resstr)
        let result = ""
    
        if (res.err){
            result = res.err
            document.getElementById("result").innerHTML = result
            return
        }
    
        res.steps.forEach(step => {
            result += "<div>"
            step.EquivalentForms.forEach(equivalentForm => {
                result += equivalentForm + "<i>→</i>"
            })
            result += step.Result + "</div>" + "<hr/>" + "<br/>"
        });
    
        document.getElementById("result").innerHTML = result
    }); 
});
