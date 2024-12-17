function search() {
    const keyword = document.getElementById("search-search-input").value;
    window.location.href = "/search?keyword=" + keyword;
}