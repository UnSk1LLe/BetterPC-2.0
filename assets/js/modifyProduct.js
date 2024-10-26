function deleteProduct(productType, productID) {
    window.location.href = `/deleteProduct?productType=${encodeURIComponent(productType)}&productID=${encodeURIComponent(productID)}`
}

function modifyProduct(productType, productID) {
    let form = document.getElementById("productForm");
    form.setAttribute("action", `/modifyProduct?productType=${encodeURIComponent(productType)}&productID=${encodeURIComponent(productID)}`);
    form.submit();
}
