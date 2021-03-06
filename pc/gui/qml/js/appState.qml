import QtQuick 2.0
import "../comps"
import "util.js" as Util

Item {
    id: appState
    property var loginSession: null
    property string path: '/'
    property var pathInfo: [{
            "path": '/',
            "name": '全部文件'
        }]
    property var enterPathPromise: null
    property var fileList: []
    property var accessDirHistory: []
    property int accessDirHistoryIndex: -1
    property var player: null
    property var mainWindow: null
    property var alertPromise: Util.Promise.resolve()
    property var pathCollection: []
    property var pathCollectionModel: null

    DataSaver {
        $key: 'app-state'
        property alias pathCollection: appState.pathCollection
    }

    onPathChanged: {
        if (path === '/') {
            pathInfo = [{
                            "path": '/',
                            "name": '全部文件'
                        }]
            return
        }
        var dirs = path.split('/')
        var arr = []
        var p = ''
        dirs.forEach(function (dir, idx) {
            p += idx === 1 ? dir : '/' + dir
            arr.push({
                         "path": p,
                         "name": idx === 0 ? '全部文件' : dir
                     })
        })
        pathInfo = arr
    }

    Component.onCompleted: {
        loginSession = {
            "username": '用户名加载中...',
            "photo": ''
        }
    }
}
