const tittle = document.title;
const video = document.getElementById('container-video')
const a = document.getElementById('linkToYTB')
const template = "https://www.youtube.com/channel/"

function iFrame(idVideo) {
    return `<iframe width="892" height="502" src="https://www.youtube.com/embed/${idVideo}" title="YouTube video player"
                   frameBorder="0"
                   allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                   allowFullScreen></iframe>`
}

fetch(`https://youtube.googleapis.com/youtube/v3/search?part=snippet&q=${tittle}&key=AIzaSyBJqDiEkpZFFtOqD9mhYiN-wRT19H6ZrqU`).then(function (response) {
    response.json().then(function (json) {
        console.log(json)

        const idChannel = json["items"][0].snippet["channelId"]
        const toDisplayVideo = json["items"][0].id["videoId"]

        video.innerHTML = iFrame(toDisplayVideo)
        a.href = template + idChannel
    })
}).catch((err) => console.log(err))
