// static/script.js
document.addEventListener('DOMContentLoaded', function() {
    // Set the selected values in dropdowns if form was submitted
    const urlParams = new URLSearchParams(window.location.search);
    const from = urlParams.get('from');
    const to = urlParams.get('to');

    if (from) {
        document.getElementById('from').value = from;
    }
    if (to) {
        document.getElementById('to').value = to;
    }

    // Prevent form submission if 'from' and 'to' units are the same
    document.querySelector('form').addEventListener('submit', function(e) {
        const fromUnit = document.getElementById('from').value;
        const toUnit = document.getElementById('to').value;

        if (fromUnit === toUnit) {
            e.preventDefault();
            alert('Please select different units for conversion');
        }
    });
});
