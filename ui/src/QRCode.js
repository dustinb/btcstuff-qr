
import React, { useRef, useState } from "react";

function QRCode() {

    const address = useRef("")
    const amount = useRef(0)
    const label = useRef("")
    const message = useRef("")

    const [fileURL, setFile] = useState("")

    async function loadQR(e) {
        e.preventDefault();
        // Validate amount
        const params = new URLSearchParams()
        params.set("address", address.current.value)
        params.set("amount", amount.current.value)
        params.set("label", label.current.value)
        params.set("message", message.current.value)

        let newURL = "/qr?" + params.toString();
        setFile(newURL)
    }

    return (
        <div className="container-fluid">
            <div className="row gy-5">
                <div className="col p-3">
                Generate a Bitcoin QR code based
                    on <a href="https://github.com/bitcoin/bips/blob/master/bip-0021.mediawiki">BIP 21</a>. Fork me
                    on <a href="https://github.com/dustinb/btcstuff-qr">Github</a>
                </div>
            </div>
            <div className="row justify-content-center">
                <div className="col">
                <form className="input-box">
                    <div className="form-group">
                        <label htmlFor="address">Receive Address</label>
                        <input ref={address} defaultValue="tb1qlmzkqda435e29s5jj74rj9jwrh0myzqw6f752x" type="text" className="form-control" id="address" aria-describedby="BTC Address" />
                    </div>
                    <div className="form-group">
                        <label htmlFor="amount">Amount</label>
                        <input ref={amount} defaultValue="" type="test" className="form-control" id="amount" aria-describedby="BTC Amount"/>
                    </div>
                    <div className="form-group">
                        <label htmlFor="label">Label</label>
                        <input ref={label} defaultValue="Testnet Address" type="text" className="form-control" id="label" aria-describedby="Label" />
                    </div>
                    <div className="form-group">
                        <label htmlFor="message">Message</label>
                        <input ref={message} defaultValue="Send test coins!" type="text" className="form-control" id="message" aria-describedby="Message" />
                    </div>
                    <button type="submit" className="btn btn-primary" onClick={(e) => loadQR(e)}>Generate QR Code</button>
                </form>
                </div>
                <div className="col">
                    {fileURL ? <div><img src={fileURL} alt="qr code" /> <br /><a href={fileURL}>Image URL</a></div> : ""}
                </div>
            </div>
        </div>
    );
}

export default QRCode;
