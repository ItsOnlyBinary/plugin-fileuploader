package templates

var Get = map[string]string{
	"Embed": `<!DOCTYPE html>
<html>
    <head>
        <style>
            html, body {
                margin: 0;
                padding: 0;
            }
            body {
                /* style.body */
            }
            .kiwi-embed-image {
                display: block;
            }
        </style>
        <script>
            let last_width = 0;
            let last_height = 0;
            window.addEventListener('load', function () {
                postDimensions();
                const observerOptions = {
                    childList: true,
                    subtree: true,
                };
                const observer = new MutationObserver(mutationObserver);
                observer.observe(document.body, observerOptions);
            })

            function mutationObserver() {
                postDimensions();
            }

            function imgError() {
                window.parent.postMessage({ error: true }, '*');
            }

            function postDimensions() {
                if (!window.parent) {
                    // Nowhere to send messages
                    return;
                }
                const div = document.body.querySelector('#kiwi-embed-container');
                const width = div.scrollWidth;
                const height = div.scrollHeight;
                if (last_width === width && last_height === height) {
                    // No change
                    return;
                }
                last_width = width;
                last_height = height;
                window.parent.postMessage({ dimensions: { width, height } }, '*');
            }
        </script>
    </head>
    <body>
        <div id="kiwi-embed-container">
            {{body.html}}
        </div>
    </body>
</html>
`,
}
