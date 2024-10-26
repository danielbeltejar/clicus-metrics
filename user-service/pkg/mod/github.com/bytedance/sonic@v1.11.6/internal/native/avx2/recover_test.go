//go:build amd64
// +build amd64

// Code generated by Makefile, DO NOT EDIT.

/**
 * Copyright 2023 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package avx2

import (
    `os`
    `runtime`
    `runtime/debug`
    `testing`
    `time`
    `unsafe`

    `github.com/bytedance/sonic/internal/native/types`
)

var (
    debugAsyncGC = os.Getenv("SONIC_NO_ASYNC_GC") == ""
)

func TestMain(m *testing.M) {
    Use()
    
    go func ()  {
        if !debugAsyncGC {
            return
        }
        println("Begin GC looping...")
        for {
        runtime.GC()
        debug.FreeOSMemory() 
        }
        println("stop GC looping!")
    }()
    time.Sleep(time.Millisecond*100)
    m.Run()
}

func TestRecover_f64toa(t *testing.T) {
    defer func() {
        if r := recover(); r!= nil {
            t.Log("recover: ", r)
        } else {
            t.Fatal("no panic")
        }
    }()
    _ = f64toa(nil, 123)
}

func TestRecover_f32toa(t *testing.T) {
    defer func() {
        if r := recover(); r!= nil {
            t.Log("recover: ", r)
        } else {
            t.Fatal("no panic")
        }
    }()
    _ = f32toa(nil, 123)
}

func TestRecover_i64toa(t *testing.T) {
    defer func() {
        if r := recover(); r!= nil {
            t.Log("recover: ", r)
        } else {
            t.Fatal("no panic")
        }
    }()
    _ = i64toa(nil, 123)
}

func TestRecover_u64toa(t *testing.T) {
    defer func() {
        if r := recover(); r!= nil {
            t.Log("recover: ", r)
        } else {
            t.Fatal("no panic")
        }
    }()
    _ = u64toa(nil, 123)
}

func TestRecover_lspace(t *testing.T) {
    defer func() {
        if r := recover(); r!= nil {
            t.Log("recover: ", r)
        } else {
            t.Fatal("no panic")
        }
    }()
    _ = lspace(nil, 2, 0)
}

func TestRecover_quote(t *testing.T) {
    var dn = 10
    var dp = make([]byte, dn)
    var sp = []byte("123")
    t.Run("sp", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = quote(nil, 3, unsafe.Pointer(&dp[0]), &dn, 0)
    })
    t.Run("dp", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = quote(unsafe.Pointer(&sp[0]), 3, nil, &dn, 0)
    })
    t.Run("dn", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = quote(unsafe.Pointer(&sp[0]), 3, unsafe.Pointer(&dp[0]), nil, 0)
    })
}

func TestRecover_html_escape(t *testing.T) {
    var dn = 10
    var dp = make([]byte, dn)
    var sp = []byte("123")
    t.Run("sp", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = html_escape(nil, 3, unsafe.Pointer(&dp[0]), &dn)
    })
    t.Run("dp", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = html_escape(unsafe.Pointer(&sp[0]), 3, nil, &dn)
    })
    t.Run("dn", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = html_escape(unsafe.Pointer(&sp[0]), 3, unsafe.Pointer(&dp[0]), nil)
    })
}

func TestRecover_unquote(t *testing.T) {
    var ep = 0
    var dp = make([]byte, 10)
    var sp = []byte("12\\x\"3\"4")
    t.Run("sp", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = unquote(nil, len(sp), unsafe.Pointer(&dp[0]), &ep, 0)
    })
    t.Run("dp", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = unquote(unsafe.Pointer(&sp[0]), len(sp), nil, &ep, 0)
    })
    t.Run("ep", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = unquote(unsafe.Pointer(&sp[0]), len(sp), unsafe.Pointer(&dp[0]), nil, 0)
    })
}

func TestRecover_value(t *testing.T) {
    var v = new(types.JsonState)
    var sp = []byte("123")
    t.Run("sp", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = value(nil, 3, 0, v, 0)
    })
    t.Run("v", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = value(unsafe.Pointer(&sp[0]), 3, 0, nil, 0)
    })
}

func TestRecover_vstring(t *testing.T) {
    var v = new(types.JsonState)
    var sp = "123"
    var p = 0
    t.Run("sp", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        vstring(nil, &p, v, 0)
    })
    t.Run("p", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        vstring(&sp, nil, v, 0)
    })
    t.Run("v", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        vstring(&sp, &p, nil, 0)
    })
}

func TestRecover_vnumber(t *testing.T) {
    var v = new(types.JsonState)
    var sp = "123"
    var p = 0
    t.Run("sp", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        vnumber(nil, &p, v)
    })
    t.Run("p", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        vnumber(&sp, nil, v)
    })
    t.Run("v", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        vnumber(&sp, &p, nil)
    })
}

func TestRecover_vsigned(t *testing.T) {
    var v = new(types.JsonState)
    var sp = "123"
    var p = 0
    t.Run("sp", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        vsigned(nil, &p, v)
    })
    t.Run("p", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        vsigned(&sp, nil, v)
    })
    t.Run("v", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        vsigned(&sp, &p, nil)
    })
}

func TestRecover_vunsigned(t *testing.T) {
    var v = new(types.JsonState)
    var sp = "123"
    var p = 0
    t.Run("sp", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        vunsigned(nil, &p, v)
    })
    t.Run("p", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        vunsigned(&sp, nil, v)
    })
    t.Run("v", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        vunsigned(&sp, &p, nil)
    })
}

func TestRecover_skip_one(t *testing.T) {
    var v = types.NewStateMachine()
    var sp = "123"
    var p = 0
    t.Run("sp", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = skip_one(nil, &p, v, 0)
    })
    t.Run("p", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = skip_one(&sp, nil, v, 0)
    })
    t.Run("v", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = skip_one(&sp, &p, nil, 0)
    })
}

func TestRecover_skip_one_fast(t *testing.T) {
    var sp = "123"
    var p = 0
    t.Run("sp", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = skip_one_fast(nil, &p)
    })
    t.Run("p", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = skip_one_fast(&sp, nil)
    })
}

func TestRecover_skip_array(t *testing.T) {
    var v = types.NewStateMachine()
    var sp = "123"
    var p = 0
    t.Run("sp", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = skip_array(nil, &p, v, 0)
    })
    t.Run("p", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = skip_array(&sp, nil, v, 0)
    })
    t.Run("v", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = skip_array(&sp, &p, nil, 0)
    })
}

func TestRecover_skip_object(t *testing.T) {
    var v = types.NewStateMachine()
    var sp = "123"
    var p = 0
    t.Run("sp", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = skip_object(nil, &p, v, 0)
    })
    t.Run("p", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = skip_object(&sp, nil, v, 0)
    })
    t.Run("v", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = skip_object(&sp, &p, nil, 0)
    })
}

func TestRecover_skip_number(t *testing.T) {
    var sp = "123"
    var p = 0
    t.Run("sp", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = skip_number(nil, &p)
    })
    t.Run("p", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = skip_number(&sp, nil)
    })
}

func TestRecover_get_by_path(t *testing.T) {
    var v = []interface{}{}
    var sp = "123"
    var p = 0
    var m = types.NewStateMachine()
    t.Run("sp", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = get_by_path(nil, &p, &v, m)
    })
    t.Run("p", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = get_by_path(&sp, nil, &v, m)
    })
    t.Run("path", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = get_by_path(&sp, &p, nil, m)
    })
}

func TestRecover_validate_one(t *testing.T) {
    var v = types.NewStateMachine()
    var sp = "123"
    var p = 0
    t.Run("sp", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = validate_one(nil, &p, v, 0)
    })
    t.Run("p", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = validate_one(&sp, nil, v, 0)
    })
    t.Run("v", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = validate_one(&sp, &p, nil, 0)
    })
}

func TestRecover_validate_utf8(t *testing.T) {
    var v = types.NewStateMachine()
    var sp = string([]byte{0xff, 0xff, 0xff})
    var p = 0
    t.Run("sp", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = validate_utf8(nil, &p, v)
    })
    t.Run("p", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = validate_utf8(&sp, nil, v)
    })
    t.Run("v", func(t *testing.T) {
        defer func() {
            if r := recover(); r!= nil {
                t.Log("recover: ", r)
            } else {
                t.Fatal("no panic")
            }
        }()
        _ = validate_utf8(&sp, &p, nil)
    })
}

func TestRecover_validate_utf8_fast(t *testing.T) {
    defer func() {
        if r := recover(); r!= nil {
            t.Log("recover: ", r)
        } else {
            t.Fatal("no panic")
        }
    }()
    _ = validate_utf8_fast(nil)
}
