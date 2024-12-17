function search() {
    const keyword = document.getElementById("index-search-input").value;
    window.location.href = "/search?keyword=" + keyword;
}