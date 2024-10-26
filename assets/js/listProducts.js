/*document.addEventListener('DOMContentLoaded', function() {
    var productData = document.getElementById('productData');
    var inCart = productData.dataset.inCart === 'true';
    var amount = parseInt(productData.dataset.amount);

    var addToCartButton = document.getElementById('addToCartButton');
    var outOfStockMessage = document.getElementById('outOfStockMessage');
    var addedText = document.getElementById('addedText');

    if (amount > 0) {
        if (inCart) {
            addToCartButton.style.display = 'none';
            addedText.style.display = 'block';
        } else {
            addToCartButton.style.display = 'block';
            addedText.style.display = 'none';

            addToCartButton.addEventListener('click', function() {
                if (amount > 0) {
                    addToCart();
                } else {
                    outOfStockMessage.style.display = 'block';
                }
            });
        }
    } else {
        outOfStockMessage.style.display = 'block';
        addToCartButton.style.display = 'none';
        addedText.style.display = 'none';
    }
});*/
let urlParams = new URLSearchParams(window.location.search);
let listCompatibleOnly = urlParams.get('listCompatibleOnly');
let search = urlParams.get('search')
listCompatibleOnly = listCompatibleOnly === 'true';

function listProducts(productType, pageNumber, searchQuery, saveFilters) {
    let urlParams = new URLSearchParams(window.location.search);
    const sort = "None"//document.getElementById('sort-price').value;
    urlParams.set('productType', productType);
    urlParams.set('listCompatibleOnly', listCompatibleOnly);
    if (saveFilters) {
        urlParams.set('pageNumber', pageNumber);
        urlParams.set('sort', sort);
        if (searchQuery) {
            urlParams.set('search', searchQuery);
        } else {
            urlParams.delete('search');
        }

        //Include current filter parameters
        document.querySelectorAll('#filters-form input, #filters-form select').forEach(input => {
            if (input.type === 'checkbox') {
                if (input.checked) {
                    let exists = false;
                    urlParams.forEach((value, name) => {
                        if (name === input.name && value === input.value) {
                            exists = true;
                        }
                    });
                    if (!exists) {
                        urlParams.append(input.name, input.value);
                    }
                } else {
                    urlParams.delete(input.name, input.value);
                }
            } else if (input.type === 'number' || input.type === 'text' || input.tagName.toLowerCase() === 'select') {
                if (input.value) {
                    let exists = false;
                    urlParams.forEach((value, name) => {
                        if (name === input.name && value === input.value) {
                            exists = true;
                        }
                    });
                    if (!exists) {
                        urlParams.set(input.name, input.value);
                    }
                } else {
                    urlParams.delete(input.name);
                }
            }
        });
    } else {
        urlParams.delete('sort');
        urlParams.delete('search');
        urlParams.set('pageNumber', '1')
    }

    window.location.href = `/listProducts?${urlParams.toString()}`;
}

/*function filterProducts(productType) {
    const form = document.getElementById('filters-form');
    form.action = `/listProducts?productType=${encodeURIComponent(productType)}`;
    form.submit();
}*/

function performSearch() {
    const searchInput = document.getElementById('search-input').value;
    let productType = urlParams.get('productType');
    listProducts(productType, 1, searchInput, true);
}

document.getElementById('search-input').addEventListener('keydown', function(event) {
    if (event.key === 'Enter') {
        performSearch();
    }
})

function modifyProductForm(productType, productID) {
    window.location.href = `/modifyProductForm?productType=${encodeURIComponent(productType)}&productID=${encodeURIComponent(productID)}`
}

function addProduct() {
    window.location.href = `/addProductForm`
}

function showProduct(productType, productID) {
    window.location.href = `/showProduct?productType=${encodeURIComponent(productType)}&productID=${encodeURIComponent(productID)}`;
}

document.addEventListener('DOMContentLoaded', function () {
    document.addEventListener('click', function (event) {
        const expandedForms = document.querySelectorAll('.component-detail-form');
        const componentContainers = document.querySelectorAll('.component');

        let isClickInsideComponent = false;

        componentContainers.forEach(container => {
            if (container.contains(event.target)) {
                isClickInsideComponent = true;
            }
        });

        expandedForms.forEach(expandedForm => {
            if (expandedForm && !expandedForm.contains(event.target) && !isClickInsideComponent) {
                expandedForm.style.display = 'none';
            }
        });
    });

});

//build functions
function handleClick(componentType) {
    const componentContainer = document.querySelector(`[data-component="${componentType}"]`);
    const isEmpty = componentContainer.querySelector('.empty-component');

    if (isEmpty) {
        listProducts(componentType);
    } else {
        expandComponent(componentType, componentContainer);
    }
}

function expandComponent(componentType, componentContainer) {
    const allDetailForms = document.querySelectorAll('.component-detail-form');
    allDetailForms.forEach(form => {
        form.style.display = 'none';
    });

    const detailFormId = `${componentType}-detail-form`;
    const detailForm = document.getElementById(detailFormId);
    if (detailForm) {
        detailForm.style.display = 'block';
    }


    detailForm.addEventListener('click', function (event) {
        event.stopPropagation();
    });

    const rect = componentContainer.getBoundingClientRect();
    detailForm.style.position = 'absolute';
    detailForm.style.top = `${rect.bottom + window.scrollY}px`;
    detailForm.style.left = `${rect.left + window.scrollX}px`;
    detailForm.style.width = `${componentContainer.offsetWidth}px`;
    detailForm.style.zIndex = 1000;
}

function replaceComponent(productType) {
    window.location.href = `/listProducts?productType=${encodeURIComponent(productType)}&listCompatibleOnly=${listCompatibleOnly}`
}

function deleteComponent(productType) {
    let xhr = new XMLHttpRequest();
    xhr.open('POST', '/deleteProductFromBuild', true);
    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE) {
            if (xhr.status === 200) {
                location.reload();
            } else {
                alert('Error deleting product from build');
            }
        }
    };
    let params = 'productType=' + encodeURIComponent(productType)
    xhr.send(params);
}

function addToCart(productType, productID, index) {
    let xhr = new XMLHttpRequest();
    xhr.open('POST', '/addProductToCart', true);
    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE) {
            if (xhr.status === 200) {
                let cartButton = document.getElementById('addToCartButton' + index)
                cartButton.style.backgroundColor = '#e3bc00';
                cartButton.innerText = 'In cart'
                cartButton.onclick = function () {
                    window.location.href = `/openCart`
                }
                cartNumberInc()
            } else {
                alert('Error adding product to cart');
            }
        }
    };
    let params = 'productType=' + encodeURIComponent(productType) + '&productID=' + encodeURIComponent(productID);
    xhr.send(params);
}

function cartNumberInc() {
    let counter = document.getElementById('cart-badge')
    let number = parseInt(counter.innerText)
    number++
    counter.innerText = number.toString()
    counter.style.display = 'block'
}

function addToBuild(productType, productID) {
    let xhr = new XMLHttpRequest();
    xhr.open('POST', '/addProductToBuild', true);
    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE) {
            if (xhr.status === 200) {
                location.reload();
            } else {
                alert('Error adding product to build');
            }
        }
    };
    let params = 'productType=' + encodeURIComponent(productType) + '&productID=' + encodeURIComponent(productID);
    xhr.send(params);
}

document.addEventListener("DOMContentLoaded", function() {
    const params = new URLSearchParams(window.location.search);
    const productType = params.get('productType');

    const filtersContainer = document.getElementById('filter-container');

    fetch(`./assets/data/filters.json`)
        .then(response => response.json())
        .then(data => {
            const filterObjectName = `${productType}Filters`;
            const filters = data[filterObjectName];
            populateFilters(filters, params);
        });

    function populateFilters(filters, params) {
        filtersContainer.innerHTML = '';
        for (const filterName in filters) {
            const filterDiv = document.createElement('div');
            filterDiv.innerHTML = `<label>${filterName}:</label><br>`;
            const options = filters[filterName];
            if (Array.isArray(options)) {
                options.forEach(option => {
                    const checked = params.getAll(filterName).includes(option.toString()) ? 'checked' : '';
                    filterDiv.innerHTML += `
                        <input type="checkbox" id="${filterName}-${option}" name="${filterName}" value="${option}" ${checked}>
                        <label for="${filterName}-${option}">${option}</label><br>
                    `;
                });
            } else if (typeof options === 'object' && options.min !== undefined && options.max !== undefined) {
                const min = params.get(`${filterName}-min`);
                const max = params.get(`${filterName}-max`);
                filterDiv.innerHTML += `
                    <label for="${filterName}-min">Min:</label>
                    <input type="number" id="${filterName}-min" name="${filterName}-min" ${min ? `value="${min}"` : `placeholder="${options.min}"`}><br>
                    <label for="${filterName}-max">Max:</label>
                    <input type="number" id="${filterName}-max" name="${filterName}-max" ${max ? `value="${max}"` : `placeholder="${options.max}"`}>
                `;
            }
            filtersContainer.appendChild(filterDiv);
        }
    }

    document.getElementById('filters-form').addEventListener('submit', function(e) {
        e.preventDefault();
        const formData = new FormData(e.target);
        const urlParams = new URLSearchParams(window.location.search);
        const productType = urlParams.get('productType');
        const pageNumber = 1; //Reset to the first page on new filter submission
        const searchQuery = urlParams.get('search');

        formData.forEach((value, key) => {
            urlParams.set(key, value);
        });

        listProducts(productType, pageNumber, searchQuery, true);
    });


    const toggleDisplay = document.getElementById('toggleDisplay');
    const build = document.getElementById('build');
    const orderBuild = document.getElementById('createOrderFromBuildButton')

    toggleDisplay.addEventListener('change', function() {
        if (this.checked) {
            build.style.display = 'flex';
            orderBuild.style.display = 'block'
            listCompatibleOnly = true
            listProducts(productType, 1, search)
        } else {
            build.style.display = 'none';
            orderBuild.style.display = 'none'
            listCompatibleOnly = false
            listProducts(productType, 1, search)
        }
    });
});

function formatPrice(price) {
    return price.toString().replace(/\B(?=(\d{3})+(?!\d))/g, " ");
}

document.querySelectorAll('.old-price').forEach(function(element) {
    let oldPrice = parseInt(element.innerText);
    element.innerText = formatPrice(oldPrice) + " ₸";
});

document.querySelectorAll('.discount-price').forEach(function(element) {
    let discountPrice = parseInt(element.innerText);
    element.innerText = formatPrice(discountPrice) + " ₸";
});

document.querySelectorAll('.product-price').forEach(function(element) {
    let productPrice = parseInt(element.innerText);
    element.innerText = formatPrice(productPrice) + " ₸";
});

function getCookie(name) {
    let value = "; " + document.cookie;
    let parts = value.split("; " + name + "=");
    if (parts.length === 2) return parts.pop().split(";").shift();
}
