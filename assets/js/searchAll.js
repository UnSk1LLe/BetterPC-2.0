let urlParams = new URLSearchParams(window.location.search);
let search = urlParams.get('search')

function searchALl() {
    const searchInput = document.getElementById('search-input').value;
    window.location.href = `/searchAll?search=${searchInput}`;
}

document.getElementById('search-input').addEventListener('keydown', function(event) {
    if (event.key === 'Enter') {
        searchALl();
    }
})

function ListProducts(productType){
    urlParams.set('productType', productType)
    urlParams.set('search', search)
    window.location.href = `/listProducts?${urlParams.toString()}`;
}