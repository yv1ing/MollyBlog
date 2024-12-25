const bodyWrap = document.getElementById('body-wrap');
const playerWrap = document.getElementById('music-player-wrap');

syncWidth();
window.addEventListener('resize', syncWidth);

function syncWidth() {
    playerWrap.style.width =window.getComputedStyle(bodyWrap).width;
}

function play(name, singer, coverSrc, musicSrc, lyricSrc) {
    fetch(lyricSrc)
        .then((res) => {
            return res.text();
        })
        .then((data) => {
            const musicData = musicSrc;
            const lyricData = data;

            MollyPlayer({
                music: {
                    src: musicData,
                    lrc: lyricData,
                    cover: coverSrc,
                    title: name,
                    author: singer,
                    loop: true,
                },
                target: '#music-player',
                autoplay: true,
            });
        });
}