function search() {
    const keyword = document.getElementById("index-search-input").value;
    window.location.href = "/search?keyword=" + keyword;
}

function typeWriter(element_id, text) {
    const element = document.getElementById(element_id);
    let i = 0;

    element.innerHTML = '';
    function type() {
        if (i < text.length) {
            const minSpeed = 50;
            const maxSpeed = 200;
            const randomSpeed = Math.floor(Math.random() * (maxSpeed - minSpeed + 1)) + minSpeed;

            element.innerHTML += text.charAt(i);

            i++;
            setTimeout(type, randomSpeed);
        }
    }

    type();
}