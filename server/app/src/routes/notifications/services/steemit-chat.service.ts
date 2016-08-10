import { Injectable }              from '@angular/core';
import { Http, Headers, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';

import { CookieService } from 'angular2-cookie/core';

import { SteemitChatModel } from '../models/steemit-chat.model';


const STEEMIT_CHAT_BASE_URL = 'https://steemit.chat';


export interface Credentials {
  userID:    string;
  authToken: string;
}


@Injectable()
export class SteemitChatService {

  constructor(
    private http: Http,
    private cookies: CookieService
  ) {}

  load() : Observable<SteemitChatModel> {
    // Send the API call.
    const url = `/api/notifiers/steemit-chat`;

    const headers = new Headers({
      'X-CSRF-Token': this.cookies.get('csrf')
    });

    return this.http.get(url, {headers})
      .map(res => {
        const payload = res.json();
        payload.enabled = payload.enabled || false;
        payload.settings = payload.settings || {};
        return <SteemitChatModel>payload;
      });
  }

  store(username: string, creds: Credentials) : Observable<Response> {
    const url = '/api/notifiers/steemit-chat';

    const headers = new Headers({
      'Content-Type': 'application/json',
      'X-CSRF-Token': this.cookies.get('csrf')
    });

    const body = JSON.stringify({
      enabled: true,
      settings: {
        username,
        userID:    creds.userID,
        authToken: creds.authToken
      }
    });

    return this.http.put(url, body, {headers});
  }

  update(model: any) : Observable<Response> {
    const url = '/api/notifiers/steemit-chat';

    const headers = new Headers({
      'Content-Type': 'application/json',
      'X-CSRF-Token': this.cookies.get('csrf')
    });

    const body = JSON.stringify(model);

    return this.http.patch(url, body, {headers});
  }

  setEnabled(enabled: boolean) : Observable<Response> {
    return this.update({enabled});
  }

  disconnect() : Observable<Response> {
    const url = '/api/notifiers/steemit-chat';

    const headers = new Headers({
      'X-CSRF-Token': this.cookies.get('csrf')
    });

    return this.http.delete(url, {headers});
  }

  logon(username: string, password: string) : Observable<Credentials> {
    const url = `${STEEMIT_CHAT_BASE_URL}/api/login`;

    const headers = new Headers({
      'Content-Type': 'application/x-www-form-urlencoded'
    });

    const body = `user=${this.encodeFormData(username)}&password=${this.encodeFormData(password)}`;

    return this.http.post(url, body, {headers})
      .map(res => {
        const data = res.json().data || {};
        if (!data.userId) {
          throw new Error('logon: userId not found');
        }
        if (!data.authToken) {
          throw new Error('logon: authToken not found');
        }
        return {
          userID:    data.userId,
          authToken: data.authToken
        };
      });
  }

  logoff(creds: Credentials) : Observable<Response> {
    const url = `${STEEMIT_CHAT_BASE_URL}/api/logoff`;

    const headers = new Headers({
      'X-User-Id':    creds.userID,
      'X-Auth-Token': creds.authToken
    });

    return this.http.get(url, {headers});
  }


  private encodeFormData(data: string) : string {
    return encodeURIComponent(data).replace('%20', '+');
  }
}
