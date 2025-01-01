const tocBox = document.getElementById('toc-box');

const tList = document.getElementById('post-content').querySelectorAll('h1, h2, h3, h4, h5, h6');
const hList = ['H1', 'H2', 'H3', 'H4', 'H5', 'H6'];

let tmpHtml = `<div>\n`;

Array.from(tList, v => {
    // get the level number
    const H = hList.indexOf(v.nodeName) + 1 || 1;
    tmpHtml += `<div class="li li-${H} toc-li"><a id="${v.id}" href="#">${v.textContent}</a></div>\n`;
});

tmpHtml += `</div>`
tocBox.innerHTML = tmpHtml;

Array.from(tList, v => {
    const btn = document.getElementById('toc-box').querySelector(`#${v.id}`);
    const ele = document.getElementById('post-content').querySelector(`#${v.id}`);

    if (!btn || !ele) {
        return;
    }

    btn.addEventListener('click', (event) => {
        event.preventDefault();
        window.scrollTo({ top: ele.offsetTop - 60, behavior: 'smooth' });
    });
});

function backToTop() {
    window.scrollTo({
        top: 0,
        behavior: 'smooth'
    })
}

// image preview
const contentImages = document.querySelectorAll("#post-content img");
contentImages.forEach((img) => {
    img.style.cursor = "pointer";
    img.addEventListener("click", () => {
        openPreview(img);
    })
});

function openPreview(img) {
    const src = img.getAttribute("src");

    document.getElementById("image-preview-img").setAttribute("src", src);
    document.getElementById("image-preview-wrap").style.display = "flex";
}

function closePreview() {
    document.getElementById("image-preview-wrap").style.display = "none";
}
