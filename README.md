<h2>ServiceQ</h2>

ServiceQ is a TCP layer for parallel service deployments. It distributes load across multiple endpoints and buffer requests on error (in scenarios of downtimes, service unavailability, connection loss etc). The buffered requests are forwarded in-order when the service is available next.

Noticeable features -

* HTTP Load Balancing<br/>
* Request retries with configurable interval<br/>
* Failed request buffering and deferred forwarding<br/>
* Concurrent connections limit<br/> 
* Add timeout to client requests
* Add headers to upstream responses
* Customizable balancer properties<br/>
* Error detection<br/>

Until I make <b>serviceq</b> available as a package download, here are the steps to run the setup - </br>

<b>Warm-up</b>

Clone the project into any directory in your workspace (say '<i>serviceq/src</i>')<br/>

<pre>$ git clone https://github.com/gptankit/serviceq/</pre>

Make sure GOPATH is pointing to <i>serviceq</i> directory<br/>
Change into directory <i>serviceq/src</i><br/>

<b>How to Build</b>

<pre>$ make ('make build' will also work)</pre>

This will create a Go binary <i>serviceq</i> in the current directory

<b>How to Install</b>

Make sure the current user has root privileges, then - </br>

<pre>$ make install</pre>

This will create a folder <i>serviceq</i> in /opt directory and copy <i>serviceq</i> binary (to <i>/opt/serviceq</i>) and <i>sq.properties</i> (load balancer configuration) file (to <i>/opt/serviceq/config</i>).<br/>

<b>How to Run</b>

Before installing, make sure the mandatory configurations in sq.properties are set (<b>LISTENER_POST</b>, <b>PROTO</b>, <b>ENDPOINTS</b>, <b>CONCURRENCY_PEAK</b>) -</br>

<pre>
#sq.properties

#Port on which serviceq listens on
LISTENER_PORT=5252

#Protocol the endpoints accept, http and https are handled differently
PROTO=http

#Endpoints seperated by comma (,) -- no spaces allowed
ENDPOINTS=http://api.server0.com:8000,http://api.server1.com:8001,http://api.server2.com:8002

#Concurrency peak defines how many max concurrent connections are allowed to the cluster
CONCURRENCY_PEAK=2048
</pre>

After all is set - </br>

<pre>$ /opt/serviceq/serviceq</pre>

Feel free to play around and post feedbacks
