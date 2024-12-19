const tagWordCloudCanvas = document.getElementById('tag-word-cloud-canvas');

let wordList = JSON.parse(tagList);

const options = {
    list: wordList,
    weightFactor: 20,
    backgroundColor: 'transparent',
    click: (item) => {
        window.location.href = "/tag/" + item[2];
    }
};

WordCloud(tagWordCloudCanvas, options);