document.addEventListener("DOMContentLoaded", function() {
    const ecores = document.getElementById('ecores');
    const ecoresBase = document.getElementById('ecoresBase');
    const ecoresBoost = document.getElementById('ecoresBoost');
    const ddr4 = document.getElementById("ddr4");
    const ddr5 = document.getElementById("ddr5");
    const ddr4input = document.getElementById("ddr4MaxFr");
    const ddr5input = document.getElementById("ddr5MaxFr");

    function setDisabledStateEcores() {
        if (ecores.value === '0') {
            ecoresBase.disabled = true;
            ecoresBoost.disabled = true;
        } else {
            ecoresBase.disabled = false;
            ecoresBoost.disabled = false;
        }

    }

    function setDisabledStateDDR() {
        ddr4input.disabled = !ddr4.checked;
        ddr5input.disabled = !ddr5.checked;
    }

    setDisabledStateEcores();
    setDisabledStateDDR()

    ecores.addEventListener('input', setDisabledStateEcores);

    // Use the 'change' event instead of 'input' for checkboxes
    ddr4.addEventListener('change', setDisabledStateDDR);
    ddr5.addEventListener('change', setDisabledStateDDR);
});
