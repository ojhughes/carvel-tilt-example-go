allow_k8s_contexts(k8s_context())
update_settings (k8s_upsert_timeout_secs = 120, max_parallel_updates=1)


port = 8084
registry = 'docker.io/ojhughes'

kapp_apply_cmd = """
    ytt --file deployments --data-value-yaml port=%d -v registry=%s| 
    kbld -f - | 
    kapp deploy -n default -a carvel-tilt-demo -y -f - > /dev/null &&
    kapp inspect -n default -a carvel-tilt-demo --raw --tty=false
""" % (port, registry)

kapp_delete_cmd = "kapp delete -n default -a carvel-tilt -y"

compile_cmd = """
    rm build/carvel-tilt || true &&
	go build \
        -o build/carvel-tilt \
        cmd/carvel-tilt/main.go &&
        echo "Go build finished\n"
"""
local_resource(
    'go-compile',
    compile_cmd,
    deps=['pkg', 'cmd', 'deployments'],
    env={'GOOS': 'linux', 'GOARCH': 'amd64'}
)

k8s_custom_deploy(
    name='carvel-tilt-demo',
    deps=['build/carvel-tilt'],
    apply_cmd=kapp_apply_cmd,
    delete_cmd=kapp_delete_cmd
)
k8s_resource('carvel-tilt-demo', port_forwards=port, auto_init=False )
