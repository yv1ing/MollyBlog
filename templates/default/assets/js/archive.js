function search() {
    const keyword = document.getElementById("archive-search-input").value;
    window.location.href = "/search?keyword=" + keyword;
}