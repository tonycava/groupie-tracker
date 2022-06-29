let range = document.getElementById('r')
let p = document.getElementById('CarrierYears2')

p.textContent = range.value

range.oninput = function () {
    p.innerHTML = "" + this.value
}

range.addEventListener("mousemove", function () {
    let x = range.value
})

let check1 = document.getElementById("1")
let check2 = document.getElementById("2")
let check3 = document.getElementById("3")
let check4 = document.getElementById("4")
let check5 = document.getElementById("5")
let check6 = document.getElementById("6")

const sub = document.getElementById("InputCredit")

sub.onclick = function () {
    if (check1.checked && check1.value === "false") {
        check1.value = "true"
    }

    if (check2.checked && check2.value === "false") {
        check2.value = "true"
    }

    if (check3.checked && check3.value === "false") {
        check3.value = "true"
    }

    if (check4.checked && check4.value === "false") {
        check4.value = "true"
    }

    if (check5.checked && check5.value === "false") {
        check5.value = "true"
    }

    if (check6.checked && check6.value === "false") {
        check6.value = "true"
    }
}


