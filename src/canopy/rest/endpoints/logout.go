/*
 * Copyright 2014 Gregory Prisament
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
package endpoints

import (
    "net/http"
    "canopy/rest/adapter"
    "canopy/rest/rest_errors"
)

func GET_POST_logout(w http.ResponseWriter, r *http.Request, info adapter.CanopyRestInfo) (map[string]interface{}, rest_errors.CanopyRestError) {
    info.Session.Values["logged_in_username"] = ""
    err := info.Session.Save(r, w)
    if err != nil {
        return nil, rest_errors.NewInternalServerError("Problem saving session")
    }

    out := map[string]interface{} {
        "result" : "ok",
    }
    return out, nil
}
