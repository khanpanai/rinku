import {useRef, useState} from 'preact/hooks'

export function App() {
    const [id, setId] = useState("")
    const [link, setLink] = useState("")
    const [error, setError] = useState("")
    const [timeoutId, setTimeoutId] = useState(0)
    const ref = useRef<HTMLInputElement>(null)
    const copiedRef = useRef<HTMLSpanElement>(null)

    const resultLink = `${window.location.origin}/${id}`
    const onSubmit = async (e: Event) => {
        e.preventDefault()
        try {
            new URL(link);
        } catch (_) {
            setError("enter valid url");
            return;
        }

        const res = await fetch("/url/create", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Accept": "application/json"
            },
            body: JSON.stringify({
                link: link
            })
        })

        if (!res.ok) {
            setError("Something went wrong");
            return;
        }

        const resJson = await res.json();

        setId(resJson.id)
        if (ref && ref.current) {
            ref.current.value = ""
        }
    }

    const onInputChange = (e: any) => {
        setError("");
        setLink((e.target as HTMLInputElement).value)
    }

    const copy = async () => {
        await navigator.clipboard.writeText(resultLink);

        clearTimeout(timeoutId)

        if (copiedRef.current) {
            copiedRef.current.classList.remove("opacity-0");
            copiedRef.current.classList.add("opacity-1");

            setTimeoutId(setTimeout(() => {
                if (copiedRef.current) {
                    copiedRef.current.classList.remove("opacity-1");
                    copiedRef.current.classList.add("opacity-0");
                }
            }, 1000))
        }
    }

    return (
        <div className="flex flex-col items-center pt-[20vh] h-screen">
            <h1 className="text-4xl font-mono">Rinku</h1>

            <h3 class="text-xl font-mono mt-3">single binary url shortener</h3>

            <div class="flex mt-3">
                <a href="" class="font-mono">short link</a>
                <span class="mx-2">·</span>
                <a href="" class="font-mono line-through" disabled>qr code</a>
            </div>

            <form onSubmit={onSubmit} class="flex flex-col w-1/2">
                <input type="text"
                       placeholder="paste link here"
                       onInput={onInputChange}
                       ref={ ref }
                       class="mt-4 py-2 px-4 text-lg outline-0 placeholder:text-lg placeholder:font-mono text-center font-mono placeholder:text-gray-400"
                />
                <span class="my-2 text-red-400 font-mono text-center">&nbsp;{ error }</span>
                <button
                    className="mt-2 font-mono bg-gray-700 text-white text-sm transition-colors hover:bg-gray-900 px-3 py-1 rounded-lg self-center select-none">Get
                    link
                </button>
            </form>

            <div class="flex flex-col gap-1">
                <button className="text-base font-mono mt-6" title="Click to copy" onClick={copy}>{ resultLink }</button>
                <span class="font-mono text-xs text-center transition-opacity opacity-0" ref={copiedRef}>copied!</span>
            </div>
        </div>
    )
}
